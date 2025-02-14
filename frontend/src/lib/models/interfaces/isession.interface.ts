import {UserModel} from "$lib/models/classes/user.model";

export interface ISession {
	name: 	string;
	token:  string;
	user: 	UserModel;
	botId: 	string;
}