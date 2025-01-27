import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { RpcError } from "@protobuf-ts/runtime-rpc";
import { useAuth } from "../state/useAuth";

export default class Api {
  private readonly transport: GrpcWebFetchTransport;
  private readonly url: string = "http://localhost:8000";

  constructor() {
    this.transport = new GrpcWebFetchTransport({
      baseUrl: this.url,
      interceptors: [
        {
          interceptUnary(next, method, input, options) {
            const auth = useAuth();

            const token = auth.state.token;

            if (token) {
              options.meta = options.meta || {};
              options.meta["Authorization"] = `Bearer ${token}`;
            }

            return next(method, input, options);
          },
        },
      ],
    });
  }

  public getGrpcWebFetchTransport(): GrpcWebFetchTransport {
    return this.transport;
  }

  public getMessageFromRequestError(
    error: unknown,
    defaultMessage: string,
  ): string {
    if (error instanceof RpcError) {
      return "[" + error.code + "] " + error.message;
    }

    console.error(error);

    return defaultMessage;
  }
}
