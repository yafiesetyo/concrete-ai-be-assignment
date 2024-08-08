export interface IAuthService {
  signUp(email: string, password: string): Promise<string>;
  signIn(email: string, password: string): Promise<string>;
  getUser(token: string): Promise<string>;
}
