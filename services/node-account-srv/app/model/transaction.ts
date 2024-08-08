import { Decimal } from "@prisma/client/runtime/library";

export interface Transaction {
  id: number;
  type: string;
  currency: string;
  amount: Decimal;
  fromAccount?: string;
  toAccount?: string;
  description?: string;
  status: string;
}
