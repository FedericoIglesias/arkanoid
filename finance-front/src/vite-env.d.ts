/// <reference types="vite/client" />

export interface Context {
  language: number;
  ID: number | null;
  jwt: string | null;
}

export interface Category {
  ID: number;
  nameCategory: string;
}

export interface MoneyType {
  ID: number;
  moneyType: string;
}

export interface getRaw {
  id: number;
  amount: number;
  description: string;
  categoryId: number;
  userId: number;
  flowId: number;
  Date: string;
  moneyTypeId: number;
}

export interface Login {
  email: string;
  password: string;
}
