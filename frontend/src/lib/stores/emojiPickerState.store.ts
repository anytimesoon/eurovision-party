import {writable} from "svelte/store";
import {CommentModel} from "$lib/models/classes/comment.model";

class EmojiPickerState {
    constructor(
        isVisible: boolean = false,
        comment: CommentModel = CommentModel.empty()
    ) {
        this.isVisible = isVisible
        this.comment = comment
    }
    public isVisible:       boolean;
    public comment:         CommentModel;
}

export const emojiPickerState = newEmojiPickerState()

function newEmojiPickerState() {
    const {subscribe, set, update} = writable(new EmojiPickerState())
    return {
        subscribe,
        set,
        update,
        close: () => set(new EmojiPickerState()),
        open: (comment: CommentModel) => openPicker(comment)
    }
}

function openPicker(comment: CommentModel) {
    if (comment.isEmpty()) {
        return
    }
    emojiPickerState.set(new EmojiPickerState(true, comment))
}