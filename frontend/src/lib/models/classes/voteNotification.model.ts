import type {CountryModel} from "$lib/models/classes/country.model";
import type {CommentModel} from "$lib/models/classes/comment.model";

export class VoteTracker {
    public count!:              number;
    public country!:            CountryModel;
    public comment!:            CommentModel;
}