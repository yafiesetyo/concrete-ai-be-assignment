import { Account } from "../model/account";
import { Login } from "../model/login";
import { User } from "../model/user";

export interface IUsecase {
  login(req: Login): Promise<string>;
  register(req: User): Promise<string>;
  createAccount(email: string, accountType: string): Promise<any>;
  getAccount(): Promise<Account[]>;
  getEmailFromToken(token: string): Promise<string>;
}
