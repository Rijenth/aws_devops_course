import { Timestamp } from "../generated/google/protobuf/timestamp";
import { UpdateUserRequest, User } from "../generated/user";

export default class UserModel implements User {
  public id: number;
  public username: string;
  public email: string;
  public firstName: string;
  public lastName: string;
  public dateOfBirth?: Timestamp;
  public address: string;
  public phoneNumber: string;
  public createdAt?: Timestamp;
  public updatedAt?: Timestamp;
  public lastLogin?: Timestamp;
  public isActive: boolean;
  public isAdmin: boolean;
  public profilePicture: string;
  public bio: string;

  constructor(user: User) {
    this.id = user.id;
    this.username = user.username;
    this.email = user.email;
    this.firstName = user.firstName;
    this.lastName = user.lastName;
    this.dateOfBirth = user.dateOfBirth;
    this.address = user.address;
    this.phoneNumber = user.phoneNumber;
    this.createdAt = user.createdAt;
    this.updatedAt = user.updatedAt;
    this.lastLogin = user.lastLogin;
    this.isActive = user.isActive;
    this.isAdmin = user.isAdmin;
    this.profilePicture = user.profilePicture;
    this.bio = user.bio;
  }

  public updateUserRequestFormat(): UpdateUserRequest {
    return {
      id: this.id,
      username: this.username,
      email: this.email,
      firstName: this.firstName,
      lastName: this.lastName,
      dateOfBirth: this.dateOfBirth,
      address: this.address,
      phoneNumber: this.phoneNumber,
      isActive: this.isActive,
      isAdmin: this.isAdmin,
      profilePicture: this.profilePicture,
      bio: this.bio,
    };
  }
}
