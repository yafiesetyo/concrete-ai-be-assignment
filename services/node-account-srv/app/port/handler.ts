import { FastifyReply, FastifyRequest } from "fastify";

export interface IHandler {
  login(req: FastifyRequest, res: FastifyReply): Promise<void>;
  register(req: FastifyRequest, res: FastifyReply): Promise<void>;
  checkSession(req: FastifyRequest, res: FastifyReply): Promise<void>;
  createAccount(req: FastifyRequest, res: FastifyReply): Promise<void>;
  getAccount(req: FastifyRequest, res: FastifyReply): Promise<void>;
}
