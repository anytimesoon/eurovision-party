import type { ISession } from "../interfaces/isession.interface";

export class ResponseModel<T> {

    public body!:    T;
    public error!:   string;

}