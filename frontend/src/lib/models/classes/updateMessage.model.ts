import type {UserModel} from "$lib/models/classes/user.model";
import type {CommentModel} from "$lib/models/classes/comment.model";

export class UpdateMessageModel {
    public updatedUser!:   UserModel;
    public comment!:       CommentModel;
}