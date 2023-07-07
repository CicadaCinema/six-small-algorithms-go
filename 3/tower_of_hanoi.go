package main

import "fmt"

const num_disks = 3

// Add a disk to the beginning of the post.
func push(post []int, disk int) []int {
	return append([]int{disk}, post...)
}

// Remove the first disk from the post.
// Return that disk and the revised post.
func pop(post []int) (int, []int) {
	return post[0], post[1:]
}

// Move one disk from from_post to to_post.
func move_disk(posts [][]int, from_post, to_post int) {
	var disk int
	disk, posts[from_post] = pop(posts[from_post])
	posts[to_post] = push(posts[to_post], disk)
}

// Draw the posts by showing the size of the disk at each level.
func draw_posts(posts [][]int) {
	// clone the posts without a reference to the originals
	temp_posts := make([][]int, num_disks)
	for i, v := range posts {
		temp_posts[i] = append([]int(nil), v...)
	}

	// add zeroes to the front of each post so they all contain num_disks disks
	for i := range temp_posts {
		og_length := len(temp_posts[i])
		for j := 0; j < num_disks-og_length; j++ {
			temp_posts[i] = push(temp_posts[i], 0)
		}
	}

	// draw each row in turn
	for row_index := 0; row_index < num_disks; row_index++ {
		for _, post := range temp_posts {
			fmt.Printf("%01d ", post[row_index])
		}
		fmt.Println()
	}
	fmt.Println("-----")
}

// Move the disks from from_post to to_post
// using temp_post as temporary storage.
func move_disks(posts [][]int, num_to_move, from_post, to_post, temp_post int) {
	// base case
	if num_to_move == 1 {
		move_disk(posts, from_post, to_post)
		return
	}

	// inductive case

	// move all but one disk to the temp post
	move_disks(posts, num_to_move-1, from_post, temp_post, to_post)

	// move the final disk to the to post
	move_disk(posts, from_post, to_post)

	// move all but one disk to the to post
	move_disks(posts, num_to_move-1, temp_post, to_post, from_post)
}

func main() {
	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := num_disks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	draw_posts(posts)

	// Move the disks.
	move_disks(posts, num_disks, 0, 1, 2)

	// Draw the final setup.
	draw_posts(posts)
}
