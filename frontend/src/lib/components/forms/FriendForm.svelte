<script lang="ts">
    import {formButtonState} from "$lib/models/enums/formButtonState.enum.js";
    import FormButton from "$lib/components/buttons/FormButton.svelte";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import {NewUserModel} from "$lib/models/classes/newUser.model";
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import {post} from "$lib/utils/genericFetch";
    import {UserModel} from "$lib/models/classes/user.model";
    import {userStore} from "$lib/stores/user.store";
    import {newUserStore} from "$lib/stores/newUser.store";

    let form: HTMLFormElement
    let formState = formButtonState.ENABLED

    const submit = async () => {
        formState = formButtonState.SENDING
        const fd = new FormData(form)
        const newUser = new NewUserModel(fd.get("name") as string)

        const createdUser: UserModel = await post<UserModel, NewUserModel>(authEP.REGISTER, newUser)
            .then(res => UserModel.deserialize(res))
        $userStore[createdUser.id] = createdUser
        $newUserStore = [...$newUserStore, newUser]

        form.reset()

        formState = formButtonState.ENABLED
    }
</script>

<form bind:this={form} on:submit|preventDefault={submit}>
    <div class="w-fit mx-auto flex justify-center">
        <input class="mr-3" id="new-user-name" type="text" name="name" placeholder="Name"/>
        <FormButton state={formState}>
            <ContentSave size="1.4em" /> Save
        </FormButton>
    </div>
</form>