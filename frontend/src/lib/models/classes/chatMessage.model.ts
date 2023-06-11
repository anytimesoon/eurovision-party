export class ChatMessageModel<T> {
    constructor(category: string, body: T) {
        this.category = category;
        this.body = body;
    }

    public category!:   string;
    public body!:   T;

}