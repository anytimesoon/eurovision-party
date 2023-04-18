<script lang="ts">
    import {sendCreateOrUpdate} from "$lib/helpers/sender.helper";
    import {NewUserModel} from "$lib/models/classes/user.model";
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import AllRegisteredUsers from "$lib/components/AllRegisteredUsers.svelte";
    import {registeredUserStore} from "$lib/stores/user.store";

    function formHandler(e:Event) {
        const formData = new FormData(e.target as HTMLFormElement)
        let newUser: NewUserModel = new NewUserModel(<string>formData.get("name"), <string>formData.get("email"))

        sendCreateOrUpdate<NewUserModel, NewUserModel>(authEP.REGISTER, newUser, "POST").then(data => {
            registeredUserStore.update( val => {
                return [...val, data.body]
            })
        })
    }
</script>

<form on:submit|preventDefault={formHandler}>
    name <input type="text" name="name" />
    email <input type="text" name="email" />
    <input type="submit">
</form>

<AllRegisteredUsers />