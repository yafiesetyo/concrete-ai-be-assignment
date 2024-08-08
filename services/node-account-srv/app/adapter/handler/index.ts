import { FastifyReply, FastifyRequest } from "fastify";
import { IHandler } from "../../port/handler";
import { IUsecase } from "../../port/usecase";
import { Login } from "../../model/login";
import { User } from "../../model/user";

export class Handler implements IHandler {
  usecase: IUsecase;

  constructor(usecase: IUsecase) {
    this.usecase = usecase;
  }

  async login(
    req: FastifyRequest<{ Body: Login }>,
    res: FastifyReply
  ): Promise<void> {
    const body = req.body;

    try {
      const token = await this.usecase.login(body);
      res.status(200).send({ message: "Ok", data: { token } });
    } catch (error: any) {
      res.status(500).send({ message: error });
    }
  }

  async checkSession(req: FastifyRequest, res: FastifyReply): Promise<void> {
    const email = req.email;
    if (!email) {
      res.status(400).send({ message: "Empty email" });
      return;
    }

    res.status(200).send({ message: "Ok", data: { email } });
  }

  async createAccount(req: FastifyRequest, res: FastifyReply): Promise<void> {
    const params = typeof req.params === "object" ? req.params : undefined;
    if (!params) {
      res.status(400).send({ message: "Empty param" });
      return;
    }

    const param: { type?: string } = params;
    const type: string = param.type || "";

    const email = req.email;
    if (!email) {
      res.status(400).send({ message: "Empty email" });
      return;
    }

    try {
      await this.usecase.createAccount(email, type);
      res.status(200).send({ message: "Ok", data: null });
    } catch (error) {
      res.status(500).send({ message: error });
    }
  }

  async getAccount(req: FastifyRequest, res: FastifyReply): Promise<void> {
    const email = req.email;
    if (!email) {
      res.status(400).send({ message: "Empty email" });
      return;
    }

    try {
      const accounts = await this.usecase.getAccount();
      res.status(200).send({ message: "Ok", data: accounts });
    } catch (error) {
      res.status(500).send({ message: error });
    }
  }

  async register(
    req: FastifyRequest<{ Body: User }>,
    res: FastifyReply
  ): Promise<void> {
    const body = req.body;

    try {
      const token = await this.usecase.register(body);
      res.status(200).send({ message: "Ok", data: { token } });
    } catch (error: any) {
      res.status(500).send({ message: error });
    }
  }
}
