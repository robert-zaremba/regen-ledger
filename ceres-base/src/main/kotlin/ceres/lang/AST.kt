package ceres.lang.ast

import ceres.data.PersistentMap
import ceres.data.avlMapOf
import ceres.lang.*
import ceres.parser.HasSourceLoc
import ceres.parser.SourceLoc

//sealed class NodeBinding {
//    data class PropValue(val name: Expr.ID, val value: Expr)
//    data class NodeID(val id: Expr.ID)
//}


// TODO SMT refinements
// TODO type inference
// TODO holes?

//data class TypeCheckEnv(val bindings: PersistentMap<String, Type>) : PersistentMap<String, Type> by bindings

data class Env(val bindings: PersistentMap<String, TypeResult.Checked> = avlMapOf()) :
    PersistentMap<String, TypeResult.Checked> by bindings {
    fun with(vararg pairs: Pair<String, TypeResult.Checked>): Env = Env(bindings.setMany(pairs.asIterable()))
    fun with(pairs: Iterable<Pair<String, TypeResult.Checked>>): Env = Env(bindings.setMany(pairs))
}

typealias EvalEnv = Env
typealias TypeCheckEnv = Env

sealed class TypeResult {
    data class Checked(val type: Type, val value: Any? = null, val hasValue: Boolean = false) : TypeResult() {
        override fun toString(): String = "${value}: ${type}"
    }

    data class Errors(val errors: List<TypeError>) : TypeResult()

    operator fun plus(err: TypeError): TypeResult.Errors =
        when (this) {
            is Checked -> TypeResult.Errors(listOf(err))
            is Errors -> TypeResult.Errors(this.errors + err)
        }

    operator fun plus(errs: TypeResult.Errors): TypeResult.Errors =
        when (this) {
            is Checked -> errs
            is Errors -> TypeResult.Errors(this.errors + errs.errors)
        }
}

data class TypeError(val msg: String, val expr: Expr?)

fun checked(type: Type, value: Any? = null, hasValue: Boolean = false): TypeResult.Checked =
    TypeResult.Checked(type, value, if (value != null) true else hasValue)

sealed class Expr : HasSourceLoc {
    abstract fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult
    protected fun error(msg: String, expr: Expr? = null): TypeResult.Errors =
        TypeResult.Errors(listOf(TypeError(msg, if (expr != null) expr else this)))
}

interface TypedFun {
    fun evalChecked(params: List<TypeResult.Checked>): TypeResult
    val type: FunctionType
}

fun checkFnCall(fnTy: FunctionType, env: Env, argsChecked: List<TypeResult.Checked>): TypeResult {
    val params = fnTy.params
    val nParams = params.size
    val envArgs = env.with(
        argsChecked.mapIndexed { idx, argTc -> params[idx].first to argTc }
    )

    fun error(msg: String): TypeResult.Errors {
        return TypeResult.Errors(listOf(TypeError(msg, null)))
    }
    for (idx in 0 until nParams) {
        val param = params[idx]
        val argTc = argsChecked[idx]
        when (val paramTc = param.second.typeCheck(envArgs, true)) {
            is TypeResult.Checked ->
                when (paramTc.type) {
                    TypeType -> {
                        val ty = paramTc.value as? Type
                        if (ty == null) {
                            return error("Can't resolve type from expression ${param.second}")
                        } else {
                            val sres = ty.checkSubType(argTc.type)
                            if (sres != null) return error(sres)
                        }
                    }
                    else -> error("param type ${paramTc.type} is of type Type")
                }
            is TypeResult.Errors -> return paramTc
        }
    }
    val retTc = fnTy.ret.typeCheck(envArgs, true)
    when (retTc) {
        is TypeResult.Checked -> {
            val retTcVal = retTc.value as? Type
            if (retTcVal == null || retTc.type !is TypeType)
                return error("return type isn't a type, got ${retTc}")
            return checked(retTcVal)
        }
        is TypeResult.Errors -> return retTc
    }
}

data class FunCall(val fn: Expr, val args: List<Expr>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        when (val fnChk = fn.typeCheck(env, eval)) {
            is TypeResult.Checked ->
                when (val fnTy = fnChk.type) {
                    is FunctionType -> {
                        val params = fnTy.params
                        val nArgs = args.size
                        val nParams = params.size
                        if (nArgs != nParams)
                            return error("Expected ${nParams} args but got ${nArgs}")
                        val argsTc = args.map { it.typeCheck(env, eval) }
                        val errs = argsTc.filterIsInstance<TypeResult.Errors>()
                        if (!errs.isEmpty())
                            return TypeResult.Errors(errs.flatMap { it.errors })
                        val argsChecked = argsTc.filterIsInstance<TypeResult.Checked>()
                        val retTc = checkFnCall(fnTy, env, argsChecked)
                        when (retTc) {
                            is TypeResult.Checked -> {
                                if (eval) {
                                    val fn = fnChk.value as? TypedFun
                                    if (fn != null)
                                        return fn.evalChecked(argsChecked)
                                }
                                return retTc
                            }
                            is TypeResult.Errors -> return retTc
                        }
                    }
                    else -> return error("Can't call non-function type ${fnTy}")
                }
            is TypeResult.Errors -> return fnChk
        }
    }
}

data class FunExpr(
    val name: String,
    val args: List<ArgDecl>,
    val retType: Expr,
    val body: Expr, override val sourceLoc: SourceLoc? = null
) : Expr() {

//    override fun reduce(env: EvalEnv): Entry {
//        args.fold(env)
//    }

    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
//        val argsChk = args.map {
//            val ty = it.type.eval(env)
//            if(ty !is Type) {
//                error("arg type doesn't evaluate to a Type", it.type)
//            } else {
//                checked(ty)
//            }
//        }
        val (env, errs) = args.fold(env to listOf<TypeResult>()) { (env, errs), arg ->
            val res = arg.type.typeCheck(env, eval)
            when (res) {

                is TypeResult.Checked ->
                    when (res.type) {
                        TypeType -> {
                            env.with(arg.name to checked(res.value as Type)) to errs
                        }
                        else -> TODO("must eval to a type")
                    }
                is TypeResult.Errors -> TODO()
            }

        }
        val res = body.typeCheck(env, eval)
        when (res) {
            is TypeResult.Checked -> {
                val ret = retType.typeCheck(env, eval)
                when (ret) {
                    is TypeResult.Checked ->
                        when (ret.type) {
                            TypeType -> {
                                val err = res.type.checkSubType(ret.value as Type)
                                TODO()
                            }
                            else -> TODO("must eval to a type")
                        }
                    is TypeResult.Errors -> TODO()
                }
            }
            is TypeResult.Errors -> TODO()
        }
    }

    inner class EvalWrapper(var env: EvalEnv, override val type: FunctionType) : TypedFun {
        override fun evalChecked(params: List<TypeResult.Checked>): TypeResult {
            if (params.size != args.size)
                throw IllegalArgumentException("${args.size} arity function invoked with ${params.size} arguments")
            env = params.foldIndexed(env) { idx, env, param ->
                env.with(args[idx].name to param)
            }
            return body.typeCheck(env, true)
        }
    }
}

data class ArgDecl(
    val name: String,
    val type: Expr,
    val defaultValue: Expr?,
    override val sourceLoc: SourceLoc? = null
) :
    HasSourceLoc

data class FunBody(val statements: List<FunStatement>)

sealed class FunStatement {
    data class LocalBinding(val name: String, val value: Expr) : FunStatement()
    data class ExprStatement(val expr: Expr) : FunStatement()
}

data class CondExpr(
    val clauses: List<Pair<Expr, Expr>>,
    val default: Expr?,
    override val sourceLoc: SourceLoc? = null
) : Expr() {
//    override fun eval(env: EvalEnv): Any? {
//        for (clause in clauses) {
//            if (clause.first.eval(env) == true)
//                return clause.second.eval(env)
//        }
//        if (default != null)
//            return default.eval(env)
//        throw IllegalStateException("Unexpected conditional expression with missing cases")
//    }

    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        var ret = clauses.fold(checked(EmptyType), { ret, clause ->
            when (val ra = clause.first.typeCheck(env, eval)) {
                is TypeResult.Checked ->
                    when (val condErr = boolType.checkSubType(ra.type)) {
                        null -> TODO()
                        else -> {
//                            ret + error(condErr, clause.first)
                            TODO()
                        }
                    }
                is TypeResult.Errors -> return ret + ra
            }
        })
        return ret
    }
}

data class TypeCaseExpr(
    val clauses: List<Pair<Expr, Expr>>,
    val default: Expr?,
    override val sourceLoc: SourceLoc? = null
) : Expr() {
    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class ObjExpr(val pairs: List<Pair<String, Expr>>, override val sourceLoc: SourceLoc? = null) : Expr() {
//    override fun eval(env: EvalEnv): Any? {
//        val kvs: MutableList<Pair<String, Any>> = mutableListOf()
//        for (kvp in pairs) {
//            var v = kvp.second.eval(env)
//            if (v != null) {
//                kvs.add(kvp.first to v)
//            }
//        }
//        return EntityImpl(avlMapOf<String, Any>(kvs))
//    }

    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

sealed class PropertyTypeExpr

data class ObjTypeExpr(val pairs: List<Pair<String, PropertyTypeExpr>>, override val sourceLoc: SourceLoc? = null) :
    Expr() {
    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class VarRef(val name: String, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult =
        when (val ty = env[name]) {
            null -> error("Unresolved references to ${name}")
            else -> ty
        }
}

data class PropertyAccess(
    val expr: Expr,
    val prop: String, override val sourceLoc: SourceLoc? = null
) : Expr() {
//    override fun eval(env: EvalEnv): Any? {
//        val obj = expr.eval(env) as Entity
//        return obj.get(prop)
//    }

    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult =
        when (val res = expr.typeCheck(env, eval)) {
            is TypeResult.Checked ->
                when (val ty = res.type) {
                    is EntityType ->
                        when (val prop = ty.properties[prop]) {
                            is OneProperty<*> -> checked(prop.type)
                            is ZeroOrOneProperty<*> -> checked(NullableType(prop.type))
                            is SetProperty<*> -> checked(SetType(prop.type))
                            is ListProperty<*> -> checked(ListType(prop.type))
                            null -> error("Property ${prop} not found in entity with type ${ty}")
                        }
                    is DisjointEntityUnion -> TODO()
                    else -> error("Can't access property on non-entity type ${ty}")
                }
            is TypeResult.Errors -> res
        }
}

data class OrExpr(val exprs: List<Expr>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

//    override fun eval(env: EvalEnv): Any? {
//        for (expr in exprs) {
//            if (expr.eval(env) == true)
//                return true
//        }
//        return false
//    }
}

data class AndExpr(val exprs: List<Expr>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

//    override fun eval(env: EvalEnv): Any? {
//        for (expr in exprs) {
//            if (expr.eval(env) == false)
//                return false
//        }
//        return true
//    }
}

data class Literal<T>(val value: T, val type: Type, override val sourceLoc: SourceLoc? = null) : Expr() {
//    override fun eval(env: EvalEnv) = value

    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult = checked(type, value, true)
}

//data class LazyExpr(val expr: Expr, override val sourceLoc: SourceLoc? = null) : Expr() {
//    override fun eval(env: EvalEnv) = TODO()
//
//    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult = TODO()
//}
//
//data class ForceExpr(val expr: Expr, override val sourceLoc: SourceLoc? = null) : Expr() {
//    override fun eval(env: EvalEnv) = TODO()
//
//    override fun typeCheck(env: TypeCheckEnv, eval: Boolean): TypeResult = TODO()
//}
//
//class LazyValue(val expr: Expr, val env: TypeCheckEnv) {
//    val value: Any? by lazy { expr.eval(env) }
//}

//data class RealL(val x: Real) : Expr()
//data class ID(val x: String) : Expr()
//data class NodeE(val props: List<NodeBinding>) : Expr()
//data class GraphE(val nodes: List<NodeE>) : Expr()
//data class Free(val x: String)
//data class SelectE(val where: List<Expr>)

enum class Visibility {
    Public, Private
}

data class Decl(
    val visibility: Visibility,
    val name: String,
    val value: Expr
)

data class Module(val decls: List<Decl>)


