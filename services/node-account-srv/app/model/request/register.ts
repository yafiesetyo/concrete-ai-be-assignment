export const RegisterRequest = {
  type: "object",
  properties: {
    email: { type: "string" },
    password: {
      type: "string",
    },
    fullname: { type: "string" },
  },
  required: ["email", "password", "fullname"],
};
