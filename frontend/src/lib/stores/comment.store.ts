import { writable } from "svelte/store";
import type { CommentModel } from "$lib/models/classes/comment.model";
import {browser} from "$app/environment";
import {CountryModel} from "$lib/models/classes/country.model";

const defaultCommentList:CommentModel[] = new Array<CommentModel>()

// export const commentStore = writable<CommentModel[]>(browser && JSON.parse(localStorage.getItem("commentStore") || JSON.stringify(defaultCommentList)));
//
// commentStore.subscribe((val) => {
//     browser && localStorage.setItem("commentStore", JSON.stringify(val))
// });

export const commentStore = writable<CommentModel[]>(defaultCommentList);