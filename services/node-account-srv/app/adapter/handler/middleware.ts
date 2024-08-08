import { FastifyReply, FastifyRequest } from "fastify";
import { IUsecase } from "../../port/usecase";

export class Middleware {
  usecase: IUsecase;

  constructor(uc: IUsecase) {
    this.usecase = uc;
  }

  async auth(req: FastifyRequest, res: FastifyReply) {
    try {
      const headers = req.headers;
      if (!headers.authorization) {
        res.status(401).send({ message: "empty header" });
        return;
      }

      const authorization = headers.authorization?.split(" ")[1];

      console.log({ authorization }, this.usecase);
      const email = await this.usecase.getEmailFromToken(authorization || "");

      req.email = email;
    } catch (error) {
      res.status(401).send({ message: "unauthorized" });
      return;
    }
  }
}
