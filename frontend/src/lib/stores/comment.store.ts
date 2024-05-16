import {derived, writable} from "svelte/store";
import type { CommentModel } from "$lib/models/classes/comment.model";

const defaultCommentList:CommentModel[] = new Array<CommentModel>()

export const commentStore = writable<CommentModel[]>(defaultCommentList);

const offset:number = 300

export const recentComments = derived(commentStore, $commentStore => {
    const currentLength = $commentStore.length
    if (currentLength < offset) {
        return $commentStore
    } else {
        return $commentStore.slice($commentStore.length - offset, $commentStore.length)
    }
})

export const olderComments = derived(commentStore, $commentStore => {
    const currentLength = $commentStore.length
    if (currentLength < offset) {
        return []
    } else {
        return $commentStore.slice(0, $commentStore.length - offset)
    }
})