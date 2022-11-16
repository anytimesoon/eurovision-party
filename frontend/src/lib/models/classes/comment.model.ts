import type { IComment } from '../interfaces/icomment.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class CommentModel implements IDeserializable<IComment>, IComment {

    public id!:         string;
	public userId!:     string;
	public text!:       string;
	public createdAt!:  Date;

	deserialize(input: IComment): this {
		Object.assign(this, input);
		return this;
	}
}