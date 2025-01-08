import {CommentModel} from "$lib/models/classes/comment.model";
import {writable} from "svelte/store";

function newReplyCommentStore() {
    const {subscribe, update, set} = writable<CommentModel>(CommentModel.empty())
    return {
        subscribe,
        update,
        set,
        close: () => set(CommentModel.empty())
    }

}

export const replyComment = newReplyCommentStore()
