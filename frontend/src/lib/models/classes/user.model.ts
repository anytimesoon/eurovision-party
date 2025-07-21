import type { IUser } from '../interfaces/iuser.interface';
import {authLvl} from "$lib/models/enums/authLvl.enum";

export class UserModel implements IUser {

    public id!:      	string;
	public name!:    	string;
	public slug!:    	string;
	public icon!:    	string;
	public authLvl!: 	number;
	public invites!: 	string[];
	public createdBy!: 	string;
	public canInvite!: 	boolean;

	constructor(
		id: string,
		name: string,
		slug: string,
		icon: string,
		authLvl: number,
		invites: string[],
		createdBy: string,
		canInvite: boolean
	) {
		this.id = id;
		this.name = name;
		this.slug = slug;
		this.icon = icon;
		this.authLvl = authLvl;
		this.invites = invites;
		this.createdBy = createdBy;
		this.canInvite = canInvite;
	}

	isAdmin():boolean {
		return this.authLvl === authLvl.ADMIN
	}

	isFriendOfFriend():boolean {
		return this.authLvl === authLvl.FRIEND_OF_FRIEND
	}

	static deserialize(input: any): UserModel {
		return new UserModel(
			input.id,
			input.name,
			input.slug,
			input.icon,
			input.authLvl,
			input.invites,
			input.createdBy,
			input.canInvite
		)
	}

	static empty(): UserModel {
		return new UserModel('', '', '', 'default', 0, [], '', false);
	}
}
