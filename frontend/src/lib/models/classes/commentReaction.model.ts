import type {ReactAction} from "$lib/models/enums/reactAction.enum";

export class CommentReactionModel {
    public action!:     ReactAction
    public userId!:     string
    public commentId!:  string
    public reaction!:   string

    constructor(action: ReactAction, userId: string, commentId: string, reaction: string) {
        this.action = action
        this.userId = userId
        this.commentId = commentId
        this.reaction = reaction
    }

    static deserialize(input: CommentReactionModel): CommentReactionModel {
        return new CommentReactionModel(input.action, input.userId, input.commentId, input.reaction)
    }
}