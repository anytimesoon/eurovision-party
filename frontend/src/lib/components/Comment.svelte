<script type="ts">
	import type { CommentModel } from "$lib/models/classes/comment.model";

    export let comment:CommentModel
</script>

<div>
    <p>{comment.userId} at {comment.createdAt}</p>
    <p>{comment.text}</p>
</div>