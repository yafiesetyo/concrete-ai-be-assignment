import { Login } from "../../model/login";
import { Transaction } from "../../model/transaction";
import { User } from "../../model/user";
import { IUsecase } from "../../port/usecase";
import { PrismaClient } from "@prisma/client";
import { Supabase } from "../supabase";
import bcrypt from "bcrypt";
import { Decimal } from "@prisma/client/runtime/library";
import { Account } from "../../model/account";

export class Usecase implements IUsecase {
  repository: PrismaClient;
  supabase: Supabase;

  constructor(supabase: Supabase) {
    this.repository = new PrismaClient();
    this.supabase = supabase;
  }

  async login(req: Login): Promise<string> {
    try {
      let token = await this.supabase.signIn(req.email, req.password);

      return token;
    } catch (error) {
      console.error("getting error when login", error);
      throw error;
    }
  }

  async register(req: User): Promise<string> {
    // check is email exists
    try {
      const current = await this.repository.users.findFirst({
        where: {
          email: req.email,
        },
      });

      if (current) {
        throw new Error("email already used");
      }

      const salt = await bcrypt.genSalt(10);
      const password = await bcrypt.hash(req.password, salt);

      // save to db
      await this.repository.users.create({
        data: {
          email: req.email,
          fullname: req.fullname,
          password,
        },
      });

      return await this.supabase.signUp(req.email, req.password);
    } catch (error) {
      console.error("getting error when login", error);
      throw error;
    }
  }

  async createAccount(email: string, accountType: string): Promise<any> {
    try {
      const user = await this.repository.users.findFirst({
        where: {
          email,
        },
      });
      console.log({ email, user });
      if (!user) {
        throw new Error("user not found");
      }

      await this.repository.accounts.create({
        data: {
          balance: new Decimal(0),
          type: accountType,
          user_id: user?.id,
          number: this.generateAccountNumber(accountType),
        },
      });
    } catch (error) {
      console.error("getting error when create account", error);
      throw error;
    }
  }

  async getAccount(): Promise<Account[]> {
    try {
      const accounts = await this.repository.accounts.findMany({
        include: {
          fromTransactions: {
            where: {
              type: "OUTCOME",
            },
          },
          toTransactions: {
            where: {
              type: "INCOME",
            },
          },
          user: true,
        },
      });

      let resp: Account[] = [];
      for (let a of accounts) {
        let fromTransactions: Transaction[] = [];
        let toTransactions: Transaction[] = [];

        console.log(a.fromTransactions, a.toTransactions);

        for (let t of a.fromTransactions) {
          fromTransactions.push({
            amount: t.amount,
            id: Number(t.id),
            type: t.type,
            currency: t.currency,
            status: t.status,
          });
        }

        for (let t of a.toTransactions) {
          toTransactions.push({
            amount: t.amount,
            id: Number(t.id),
            type: t.type,
            currency: t.currency,
            status: t.status,
          });
        }

        resp.push({
          id: Number(a.id),
          number: a.number!,
          type: a.type,
          balance: a.balance,
          email: a.user.email,
          fullname: a.user.fullname,
          inTransaction: toTransactions,
          outTransaction: fromTransactions,
        });
      }

      return resp;
    } catch (error) {
      throw error;
    }
  }

  private generateAccountNumber(accountType: string): string {
    let prefix = "AAA";
    switch (accountType.toLowerCase()) {
      case "credit":
        prefix = "BBB";
      case "debit":
        prefix = "CCC";
      case "loan":
        prefix = "DDD";
    }

    const now = Date.now().toString();

    return `${prefix}${now}`;
  }

  async getEmailFromToken(token: string): Promise<string> {
    try {
      const email = await this.supabase.getUser(token);
      return email;
    } catch (error) {
      console.error(error);
      throw error;
    }
  }
}
