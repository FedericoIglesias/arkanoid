/// <reference types="vite/client" />

export interface Context {
  language: number;
  userId: number;
}

export interface Category {
  ID: number;
  nameCategory: string;
}

export interface MoneyType {
  ID: number;
  moneyType: string;
}

interface getRaw {
  id: number;
  amount: number;
  description: string;
  categoryId: number;
  userId: number;
  flowId: number;
  Date: string;
  moneyTypeId: number;
}
