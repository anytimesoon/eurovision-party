import type { IUser } from '../interfaces/iuser.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class UserModel implements IDeserializable<IUser>, IUser {

    public id!:   string;
	public name!: string;
	public slug!: string;
	public icon!: string;

	deserialize(input: IUser): this {
		Object.assign(this, input);
		return this;
	}
}