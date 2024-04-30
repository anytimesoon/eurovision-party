import {CommentModel} from "$lib/models/classes/comment.model";
import {writable} from "svelte/store";

function newReplyCommentStore() {
    const {subscribe, update, set} = writable<CommentModel>(new CommentModel())
    return {
        subscribe,
        update,
        set,
        close: () => set(new CommentModel())
    }

}

export const replyComment = newReplyCommentStore()
// export const replyComment = writable<CommentModel>(new CommentModel())
