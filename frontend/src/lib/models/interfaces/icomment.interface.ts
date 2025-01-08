export interface IComment {
    id:         	string;
	userId:     	string;
	text:       	string;
	createdAt:  	Date;
	isCompact:		boolean;
	replyToComment: IComment;
	fileName: string;

	isEmpty(): boolean;
}