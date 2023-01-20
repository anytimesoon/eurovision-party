import type { IUser } from '../interfaces/iuser.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class UserModel implements IDeserializable<IUser>, IUser {

    public id!:      string;
	public name!:    string;
	public slug!:    string;
	public icon!:    string;
	public authLvl!: number;

	deserialize(input: IUser): this {
		Object.assign(this, input);
		return this;
	}
}

export class NewUserModel {
	public id!:      string;
	public name!:    string;
	public email!:   string;
	public slug!:    string;
	public token!:   string;


	constructor(name: string, email: string) {
		this.name = name;
		this.email = email;
	}
}