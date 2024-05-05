import type { IComment } from '../interfaces/icomment.interface'
import type {IDeserializable} from "$lib/models/interfaces/ideserializable.interface";
import {v4 as uuid} from 'uuid';

export class CommentModel implements IComment, IDeserializable<string> {
    constructor(text?: string,
				userId?: string,
				replyToComment?: CommentModel,
				createdAt?: Date,
				isCompact = false,
				fileName = "") {
		this.id = uuid()
        this.text = text
		this.userId = userId
		this.createdAt = createdAt
		if (replyToComment && replyToComment.id != undefined) {
			this.replyToComment = replyToComment
		}
		this.isCompact = isCompact
		this.fileName = fileName
    }


    public id!:          		string;
	public userId!:      		string;
	public text!:        		string;
	public createdAt!:   		Date;
	public isCompact:	 		boolean;
	public replyToComment: 		CommentModel;
	public fileName:			string;


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