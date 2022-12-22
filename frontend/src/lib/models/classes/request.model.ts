export class RequestModel<T> {

    public token!:   string;
    public body!:    T;

    public build = (token:string, body:any):this => {
        this.token = token;
        this.body = body;

        return this
    }
}