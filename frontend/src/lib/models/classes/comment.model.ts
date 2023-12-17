import type { IComment } from '../interfaces/icomment.interface'
import type {IDeserializable} from "$lib/models/interfaces/ideserializable.interface";

export class CommentModel implements IComment, IDeserializable<string> {
    constructor(text?: string, userId?: string, replyToComment?: CommentModel, createdAt?: Date) {
        this.text = text
		this.userId = userId
		this.createdAt = createdAt
		if (replyToComment && replyToComment.id != undefined) {
			this.replyToComment = replyToComment
		}
    }


    public id!:          		string;
	public userId!:      		string;
	public text!:        		string;
	public createdAt!:   		Date;
	public isCompact:	 		boolean = false;
	public replyToComment: 		CommentModel;


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