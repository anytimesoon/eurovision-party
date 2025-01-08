import type { IVote } from '../interfaces/ivote.interface';

export class VoteModel implements IVote {

    public id!:           string;
	public userId!:       string;
	public countrySlug!:  string;
	public costume!:      number;
	public song!:         number;
	public performance!:  number;
	public props!:        number;

	constructor(
		id: string,
		userId: string,
		countrySlug: string,
		costume: number,
		song: number,
		performance: number,
		props: number)
	{
		this.id = id
		this.userId = userId
		this.countrySlug = countrySlug
		this.costume = costume
		this.song = song
		this.performance = performance
		this.props = props

	}

	static deserialize(input: IVote): VoteModel {
		return new VoteModel(
			input.id,
			input.userId,
			input.countrySlug,
			input.costume,
			input.song,
			input.performance,
			input.props
		);
	}

	static empty(): IVote {
		return new VoteModel(
			"",
			"",
			"",
			0,
			0,
			0,
			0,
		);
	}
}