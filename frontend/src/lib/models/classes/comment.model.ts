import type {IComment} from '../interfaces/icomment.interface'
import {SvelteMap, SvelteSet} from "svelte/reactivity";
import {ReactAction} from "$lib/models/enums/reactAction.enum";

export class CommentModel implements IComment {
	constructor(
		text: string,
		userId: string,
		id?: string,
		replyToComment?: CommentModel,
		createdAt?: Date,
		isCompact: boolean = false,
		fileName: string = "",
		isVoteNotification: boolean = false,
		reactions: SvelteMap<string, SvelteSet<string>> = new SvelteMap<string, SvelteSet<string>>()
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
		this.isVoteNotification = isVoteNotification
		this.reactions = reactions
	}


    public id:          		string;
	public userId!:      		string;
	public text!:        		string;
	public createdAt!:   		Date;
	public isCompact:	 		boolean;
	public replyToComment: 		CommentModel;
	public fileName:			string;
	public isVoteNotification:	boolean;
	public reactions: 			SvelteMap<string, SvelteSet<string>>;

	isEmpty(): boolean {
		return this.text === ""
			&& this.userId === ""
			&& this.id === null
			&& this.replyToComment === undefined
			&& this.isCompact === false
			&& this.fileName === ""
			&& this.reactions.size === 0
	}

	static deserialize(input: IComment): CommentModel {
        const sMap = new SvelteMap<string, SvelteSet<string>>()
        for (const [key, value] of Object.entries(input.reactions)) {
            sMap.set(key, new SvelteSet(value))
        }
		return new CommentModel(
			input.text,
			input.userId,
			input.id,
			input.replyToComment,
			new Date(input.createdAt),
			input.isCompact,
			input.fileName,
			input.isVoteNotification,
			sMap
		)
	}

	static empty(): CommentModel {
		return new CommentModel(
			"",
			"",
			null,
			undefined,
			new Date(),
			false,
			"",
			false,
			new SvelteMap<string, SvelteSet<string>>()
		)
	}

	addOrRemoveReaction = (userId: string, reaction: string): ReactAction => {
		if (this.reactions.get(reaction)?.has(userId)) {
            this.reactions.get(reaction).delete(userId)
            if (this.reactions.get(reaction).size === 0) {
				this.reactions.delete(reaction)
			}
            return ReactAction.DELETE
		} else if (this.reactions.get(reaction)) {
            this.reactions.get(reaction).add(userId)
            return ReactAction.ADD
		} else {
            this.reactions.set(reaction, new SvelteSet([userId]))
            return ReactAction.ADD
        }
	}
}