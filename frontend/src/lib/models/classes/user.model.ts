import type { IUser } from '../interfaces/iuser.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';
import {authLvl} from "$lib/models/enums/authLvl.enum";

export class UserModel implements IDeserializable<IUser>, IUser {

    public id!:      string;
	public name!:    string;
	public slug!:    string;
	public icon!:    string;
	public authLvl!: number;

	isAdmin():boolean {
		return this.authLvl === authLvl.ADMIN
	}

	deserialize(input: IUser): this {
		Object.assign(this, input);
		return this;
	}
}

export class NewUserModel {
	public id!:      string;
	public name!:    string;
	public slug!:    string;
	public token!:   string;


	constructor(name: string) {
		this.name = name;
	}
}