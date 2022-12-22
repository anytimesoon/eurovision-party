import type { IToken } from "../interfaces/itoken.interface";

export class ResponseModel<T> {

    public token!:   IToken;
    public body!:    T;
    public error!:   string;

}