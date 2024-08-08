import { IAuthService } from "../../port/auth";
import { createClient, SupabaseClient } from "@supabase/supabase-js";

export class Supabase implements IAuthService {
  client: SupabaseClient;
  constructor(url: string, key: string) {
    this.client = createClient(url, key);
  }

  async signUp(email: string, password: string): Promise<string> {
    const { data, error } = await this.client.auth.signUp({
      email,
      password,
    });

    if (error) {
      throw new Error(error.message);
    }

    const accessToken: string | undefined = data.session?.access_token;

    if (!accessToken) {
      return "";
    }

    return accessToken!;
  }

  async signIn(email: string, password: string): Promise<string> {
    const { data, error } = await this.client.auth.signInWithPassword({
      email,
      password,
    });

    if (error) {
      throw error;
    }

    const accessToken: string | undefined = data.session?.access_token;

    if (!accessToken) {
      return "";
    }

    return accessToken!;
  }

  async getUser(token: string): Promise<string> {
    const { data, error } = await this.client.auth.getUser(token);

    if (error) {
      throw error;
    }

    if (!data) {
      throw "invalid user";
    }

    return data.user.email || "";
  }
}
