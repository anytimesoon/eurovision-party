import type { IComment } from '../interfaces/icomment.interface'
import type {IDeserializable} from "$lib/models/interfaces/ideserializable.interface";

export class CommentModel implements IComment, IDeserializable<string> {
    constructor(text?: string, userId?: string) {
        this.text = text;
		this.userId = userId;
    }


    public id!:          string;
	public userId!:      string;
	public text!:        string;
	public createdAt!:   Date;

	deserialize(input: string): this {
		const obj = JSON.parse(input, function reviver(key, value) {
			if (typeof value === "string" && key === "createdAt") {
				return new Date(value);
			}

			return value;
		});
		Object.assign(this, obj);
		return this;
	}
}