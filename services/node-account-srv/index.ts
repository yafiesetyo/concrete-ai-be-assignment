import fastify, { FastifyReply, FastifyRequest } from "fastify";
import { Handler } from "./app/adapter/handler";
import { Usecase } from "./app/adapter/usecase";
import { Supabase } from "./app/adapter/supabase";
import { LoginRequest } from "./app/model/request/login";
import { RegisterRequest } from "./app/model/request/register";
import { Middleware } from "./app/adapter/handler/middleware";

declare module "fastify" {
  interface FastifyRequest {
    email?: string;
  }
}

const supabase = new Supabase(
  "https://anjpqqkwcjrezavtcosm.supabase.co",
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFuanBxcWt3Y2pyZXphdnRjb3NtIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjI5NTYwOTgsImV4cCI6MjAzODUzMjA5OH0.ORm5Vgcyepi7BAYZzYWG_S-1OXgNaR2Q1DlppsmfAwI"
);
const usecase = new Usecase(supabase);
const handlers = new Handler(usecase);
const middleware = new Middleware(usecase);

const server = fastify({
  logger: true,
});

const port = process.env.PORT || "8000";

server.post(
  "/accounts/login",
  {
    schema: {
      body: LoginRequest,
    },
  },
  handlers.login.bind(handlers)
);

server.post(
  "/accounts/register",
  {
    schema: {
      body: RegisterRequest,
    },
  },
  handlers.register.bind(handlers)
);

server.post(
  "/accounts/create/:type",
  {
    onRequest: [middleware.auth.bind(middleware)],
  },
  handlers.createAccount.bind(handlers)
);

server.get(
  "/accounts",
  {
    onRequest: [middleware.auth.bind(middleware)],
  },
  handlers.getAccount.bind(handlers)
);

server.get(
  "/accounts/session",
  {
    onRequest: [middleware.auth.bind(middleware)],
  },
  handlers.checkSession.bind(handlers)
);

server.listen({ port: Number(port), host: "0.0.0.0" }, (err, address) => {
  if (err) {
    console.error(err);
    process.exit(1);
  }

  console.log(`Server listening at ${address}`);
});
