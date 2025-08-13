import {derived, writable} from "svelte/store";
import type { CommentModel } from "$lib/models/classes/comment.model";

const defaultCommentList = new Map<string, CommentModel>()

export const commentStore = writable<Map<string, CommentModel>>(defaultCommentList);

const offset:number = 300

export const recentComments = derived(commentStore, $commentStore => {
    const currentLength = $commentStore.size
    if (currentLength < offset) {
        return $commentStore.values().toArray()
    } else {
        return $commentStore.values().toArray().slice($commentStore.size - offset, $commentStore.size)
    }
})

export const olderComments = derived(commentStore, $commentStore => {
    const currentLength = $commentStore.size
    if (currentLength < offset) {
        return []
    } else {
        return $commentStore.values().toArray().slice(0, $commentStore.size - offset)
    }
})