<script lang="ts">
    import FormButton from "$lib/components/buttons/FormButton.svelte";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {put} from "$lib/utils/genericFetch";
    import {UserModel} from "$lib/models/classes/user.model";
    import {userEP} from "$lib/models/enums/endpoints.enum";

    export let formToggle: VoidFunction

    let formState = formButtonState.ENABLED
    let form: HTMLFormElement

    const submit = async () => {
        formState = formButtonState.SENDING;

        const formData = new FormData(form);
        const newUser = new UserModel(
            $currentUser.id,
            formData.get("name") as string,
            $currentUser.slug,
            $currentUser.icon,
            $currentUser.authLvl
        )

        const updatedUser = await put<UserModel, UserModel>(userEP.UPDATE, newUser)
            .then(res => UserModel.deserialize(res))
        $currentUser = updatedUser
        $userStore[updatedUser.id] = updatedUser

        form.reset()
        formToggle()

        formState = formButtonState.ENABLED
    }
</script>
<form on:submit|preventDefault={submit} bind:this={form}>
    <div class="w-fit mx-auto flex justify-center">
        <input class="mr-3" type="text" name="name" bind:value={$currentUser.name} placeholder="Change your name"/>
        <FormButton state={formState}>
            <ContentSave size="1.4em" />
        </FormButton>
    </div>
</form>