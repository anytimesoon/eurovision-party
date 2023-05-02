import type { IVote } from '../interfaces/ivote.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class VoteModel implements IDeserializable<IVote>, IVote {

    public id!:           string;
	public userId!:       string;
	public countrySlug!:  string;
	public costume!:      number;
	public song!:         number;
	public performance!:  number;
	public props!:        number;

	deserialize(input: IVote): this {
		Object.assign(this, input);
		return this;
	}
}