import {
  DeleteUserResponse,
  GetAllUsersResponse,
  UpdateUserRequest,
  UserResponse,
} from "../generated/user";
import { UserServiceClient } from "../generated/user.client";
import Api from "./api";

export default class UserApi extends Api {
  private readonly userServiceClient: UserServiceClient;

  constructor() {
    super();
    this.userServiceClient = new UserServiceClient(
      this.getGrpcWebFetchTransport(),
    );
  }

  public async index(): Promise<GetAllUsersResponse | string> {
    try {
      const request = await this.userServiceClient.index({});

      const users = request.response?.users || [];

      return { users };
    } catch (error) {
      return this.getMessageFromRequestError(
        error,
        "[user.api.ts] Erreur lors de la récupération des utilisateurs",
      );
    }
  }

  public async getByUsername(username: string): Promise<UserResponse | string> {
    try {
      const request = await this.userServiceClient.get({ username });

      const user = request.response?.user || undefined;

      return { user };
    } catch (error) {
      return this.getMessageFromRequestError(
        error,
        "[user.api.ts] Erreur lors de la récupération de l'utilisateur",
      );
    }
  }

  public async patch(
    updateUserRequest: UpdateUserRequest,
  ): Promise<UserResponse | string> {
    try {
      const request = await this.userServiceClient.patch(updateUserRequest);

      const user = request.response?.user || undefined;

      return { user };
    } catch (error) {
      return this.getMessageFromRequestError(
        error,
        "[user.api.ts] Erreur lors de la mise à jour de l'utilisateur",
      );
    }
  }

  public async delete(id: number): Promise<DeleteUserResponse | string> {
    try {
      const request = await this.userServiceClient.delete({ id });

      return request.response;
    } catch (error) {
      return this.getMessageFromRequestError(
        error,
        "[user.api.ts] Erreur lors de la suppression de l'utilisateur",
      );
    }
  }
}
