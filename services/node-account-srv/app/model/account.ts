import { Decimal } from "@prisma/client/runtime/library";
import { Transaction } from "./transaction";

export interface Account {
  id: number;
  number: string;
  type: string;
  balance: Decimal;
  email: string;
  fullname: string;
  transactions?: Transaction[];

  inTransaction?: Transaction[];
  outTransaction?: Transaction[];
}
