package GKEgRPCService

import GKEgRPCService.GKEgRPCServiceGrpc.*

import io.grpc.*
import io.grpc.stub.*
import io.rouz.grpc.*

import kotlin.coroutines.*
import kotlinx.coroutines.*
import kotlinx.coroutines.channels.*



@javax.annotation.Generated(
    value = ["by gRPC Kotlin generator"],
    comments = "Source: interface.proto"
)
abstract class GKEgRPCServiceImplBase(
    coroutineContext: CoroutineContext = Dispatchers.Default
) : BindableService, CoroutineScope {

    private val _coroutineContext: CoroutineContext = coroutineContext

    override val coroutineContext: CoroutineContext
        get() = ContextCoroutineContextElement() + _coroutineContext

    
    
    
    open suspend fun test(request: GKEgRPCService.Interface.TestRequest): GKEgRPCService.Interface.TestResponse {
        throw unimplemented(getTestMethod()).asRuntimeException()
    }

    internal fun testInternal(
        request: GKEgRPCService.Interface.TestRequest,
        responseObserver: StreamObserver<GKEgRPCService.Interface.TestResponse>
    ) {
        launch {
            tryCatchingStatus(responseObserver) {
                val response = test(request)
                onNext(response)
            }
        }
    }
    
    
    
    open suspend fun helloWorld(request: GKEgRPCService.Interface.HelloWorldRequest): GKEgRPCService.Interface.HelloWorldResponse {
        throw unimplemented(getHelloWorldMethod()).asRuntimeException()
    }

    internal fun helloWorldInternal(
        request: GKEgRPCService.Interface.HelloWorldRequest,
        responseObserver: StreamObserver<GKEgRPCService.Interface.HelloWorldResponse>
    ) {
        launch {
            tryCatchingStatus(responseObserver) {
                val response = helloWorld(request)
                onNext(response)
            }
        }
    }

    override fun bindService(): ServerServiceDefinition {
        return ServerServiceDefinition.builder(getServiceDescriptor())
            .addMethod(
                getTestMethod(),
                ServerCalls.asyncUnaryCall(
                    MethodHandlers(METHODID_TEST)
                )
            )
            .addMethod(
                getHelloWorldMethod(),
                ServerCalls.asyncUnaryCall(
                    MethodHandlers(METHODID_HELLO_WORLD)
                )
            )
            .build()
    }

    private fun unimplemented(methodDescriptor: MethodDescriptor<*, *>): Status {
        return Status.UNIMPLEMENTED
            .withDescription("Method ${methodDescriptor.fullMethodName} is unimplemented")
    }

    private fun <E> handleException(t: Throwable?, responseObserver: StreamObserver<E>) {
        when (t) {
            null -> return
            is CancellationException -> handleException(t.cause, responseObserver)
            is StatusException, is StatusRuntimeException -> responseObserver.onError(t)
            is RuntimeException -> {
                responseObserver.onError(Status.UNKNOWN.asRuntimeException())
                throw t
            }
            is Exception -> {
                responseObserver.onError(Status.UNKNOWN.asException())
                throw t
            }
            else -> {
                responseObserver.onError(Status.INTERNAL.asException())
                throw t
            }
        }
    }

    private suspend fun <E> tryCatchingStatus(responseObserver: StreamObserver<E>, body: suspend StreamObserver<E>.() -> Unit) {
        try {
            responseObserver.body()
            responseObserver.onCompleted()
        } catch (t: Throwable) {
            handleException(t, responseObserver)
        }
    }

    private val METHODID_TEST = 0
    private val METHODID_HELLO_WORLD = 1

    private inner class MethodHandlers<Req, Resp> internal constructor(
        private val methodId: Int
    ) : ServerCalls.UnaryMethod<Req, Resp>,
        ServerCalls.ServerStreamingMethod<Req, Resp>,
        ServerCalls.ClientStreamingMethod<Req, Resp>,
        ServerCalls.BidiStreamingMethod<Req, Resp> {

        @Suppress("UNCHECKED_CAST")
        override fun invoke(request: Req, responseObserver: StreamObserver<Resp>) {
            when (methodId) {
                METHODID_TEST ->
                    this@GKEgRPCServiceImplBase.testInternal(
                        request as GKEgRPCService.Interface.TestRequest,
                        responseObserver as StreamObserver<GKEgRPCService.Interface.TestResponse>
                    )
                METHODID_HELLO_WORLD ->
                    this@GKEgRPCServiceImplBase.helloWorldInternal(
                        request as GKEgRPCService.Interface.HelloWorldRequest,
                        responseObserver as StreamObserver<GKEgRPCService.Interface.HelloWorldResponse>
                    )
                else -> throw AssertionError()
            }
        }

        @Suppress("UNCHECKED_CAST")
        override fun invoke(responseObserver: StreamObserver<Resp>): StreamObserver<Req> {
            when (methodId) {
                else -> throw AssertionError()
            }
        }
    }
}
