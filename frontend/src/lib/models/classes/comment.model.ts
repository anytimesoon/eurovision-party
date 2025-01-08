import type { IComment } from '../interfaces/icomment.interface'

export class CommentModel implements IComment {
	constructor(
		text: string,
		userId: string,
		id?: string,
		replyToComment?: CommentModel,
		createdAt?: Date,
		isCompact: boolean = false,
		fileName: string = ""
	) {
		this.id = id
		this.text = text
		this.userId = userId
		this.createdAt = createdAt
		if (replyToComment && replyToComment.id != undefined) {
			this.replyToComment = replyToComment
		}
		this.isCompact = isCompact
		this.fileName = fileName
	}


    public id:          		string;
	public userId!:      		string;
	public text!:        		string;
	public createdAt!:   		Date;
	public isCompact:	 		boolean;
	public replyToComment: 		CommentModel;
	public fileName:			string;

	isEmpty(): boolean {
		console.log(this.text === ""
			&& this.userId === ""
			&& this.id === null
			&& this.replyToComment === undefined
			&& this.isCompact === false
			&& this.fileName === "")
		return this.text === ""
			&& this.userId === ""
			&& this.id === null
			&& this.replyToComment === undefined
			&& this.isCompact === false
			&& this.fileName === ""
	}

	static deserialize(input: IComment): CommentModel {
		return new CommentModel(
			input.text,
			input.userId,
			input.id,
			input.replyToComment,
			new Date(input.createdAt),
			input.isCompact,
			input.fileName
		)
	}

	static empty(): CommentModel {
		return new CommentModel(
			"",
			"empty",
			null,
			undefined,
			new Date(),
			false,
			""
		)
	}
}