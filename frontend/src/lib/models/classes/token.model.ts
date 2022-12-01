import type { IToken } from '../interfaces/itoken.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class TokenModel implements IDeserializable<IToken>, IToken {

    public token!:   string;

	deserialize(input: IToken): this {
		Object.assign(this, input);
		return this;
	}
}