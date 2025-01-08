import type { IUser } from '../interfaces/iuser.interface';
import {authLvl} from "$lib/models/enums/authLvl.enum";

export class UserModel implements IUser {

    public id!:      string;
	public name!:    string;
	public slug!:    string;
	public icon!:    string;
	public authLvl!: number;

	constructor(id: string, name: string, slug: string, icon: string, authLvl: number) {
		this.id = id;
		this.name = name;
		this.slug = slug;
		this.icon = icon;
		this.authLvl = authLvl;
	}

	isAdmin():boolean {
		return this.authLvl === authLvl.ADMIN
	}

	static deserialize(input: any): UserModel {
		return new UserModel(
			input.id,
			input.name,
			input.slug,
			input.icon,
			input.authLvl
		)
	}

	static empty(): UserModel {
		return new UserModel('', '', '', 'default', 0);
	}
}
