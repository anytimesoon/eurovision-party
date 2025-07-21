export interface IUser {
	id:   		string;
	name: 		string;
	slug: 		string;
	icon: 		string;
	authLvl: 	number;
	invites: 	string[];
	createdBy: 	string;
	canInvite: 	boolean;
}