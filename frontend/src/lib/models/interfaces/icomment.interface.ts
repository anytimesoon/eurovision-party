import {SvelteMap, type SvelteSet} from "svelte/reactivity";
import type {ReactAction} from "$lib/models/enums/reactAction.enum";

export interface IComment {
    id:         		string;
	userId:     		string;
	text:       		string;
	createdAt:  		Date;
	isCompact:			boolean;
	replyToComment: 	IComment;
	fileName: 			string;
	isVoteNotification: boolean;
	reactions: 			SvelteMap<string, SvelteSet<string>>;

	isEmpty(): boolean;
	addOrRemoveReaction(userId: string, reaction: string): ReactAction;
}