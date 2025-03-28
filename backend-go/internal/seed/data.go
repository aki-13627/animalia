package seed

import (
	"github.com/google/uuid"
	"github.com/htanos/animalia/backend-go/internal/domain/models"
)

// UserData returns a list of sample users
func UserData() []models.User {
	return []models.User{
		{
			ID:           uuid.New().String(),
			Email:        "john.doe@example.com",
			Name:         "John Doe",
			Bio:          "I'm a pet shop owner",
			IconImageKey: "",
		},
		{
			ID:           uuid.New().String(),
			Email:        "jane.smith@example.com",
			Name:         "Jane Smith",
			Bio:          "I'm a cat lover",
			IconImageKey: "",
		},
		{
			ID:           uuid.New().String(),
			Email:        "alex.johnson@example.com",
			Name:         "Alex Johnson",
			Bio:          "I'm a dog lover",
			IconImageKey: "",
		},
		{
			ID:           uuid.New().String(),
			Email:        "emily.wilson@example.com",
			Name:         "Emily Wilson",
			Bio:          "I'm a food lover",
			IconImageKey: "",
		},
		{
			ID:           uuid.New().String(),
			Email:        "michael.brown@example.com",
			Name:         "Michael Brown",
			Bio:          "I'm a flower shop owner",
			IconImageKey: "",
		},
		{
			ID:           uuid.New().String(),
			Email:        "tanomitsu2002@gmail.com",
			Name:         "Mitsuru Hatano",
			Bio:          "I'm a software engineer",
			IconImageKey: "",
		},
		{
			ID:           uuid.New().String(),
			Email:        "aki.kaku0627@gmail.com",
			Name:         "Akihiro Kaku",
			Bio:          "I'm a software engineer",
			IconImageKey: "",
		},
	}
}

// PetData returns a list of sample pets
func PetData(users []models.User) []models.Pet {
	return []models.Pet{
		{
			ID:       uuid.New().String(),
			Name:     "Max",
			BirthDay: "2023-01-15",
			Type:     "dog",
			Species:  "saluki",
			ImageKey: "pets/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
			OwnerID:  users[0].ID, // John's pet
		},
		{
			ID:       uuid.New().String(),
			Name:     "Luna",
			BirthDay: "2022-05-10",
			Type:     "cat",
			Species:  "siamese",
			ImageKey: "pets/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
			OwnerID:  users[1].ID, // Jane's pet
		},
		{
			ID:       uuid.New().String(),
			Name:     "Buddy",
			BirthDay: "2021-11-22",
			Type:     "dog",
			Species:  "beagle",
			ImageKey: "pets/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
			OwnerID:  users[2].ID, // Alex's pet
		},
		{
			ID:       uuid.New().String(),
			Name:     "Coco",
			BirthDay: "2023-03-05",
			Type:     "dog",
			Species:  "poodle",
			ImageKey: "pets/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
			OwnerID:  users[3].ID, // Emily's pet
		},
		{
			ID:       uuid.New().String(),
			Name:     "Rocky",
			BirthDay: "2022-08-17",
			Type:     "dog",
			Species:  "golden_retriever",
			ImageKey: "pets/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
			OwnerID:  users[4].ID, // Michael's pet
		},
		{
			ID:       uuid.New().String(),
			Name:     "Milo",
			BirthDay: "2023-02-28",
			Type:     "cat",
			Species:  "munchkin",
			ImageKey: "pets/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
			OwnerID:  users[0].ID, // John's second pet
		},
	}
}

// PostData returns a list of sample posts
func PostData(users []models.User) []models.Post {
	return []models.Post{
		{
			ID:       uuid.New().String(),
			Caption:  "Max's first day at the park",
			UserID:   users[0].ID, // John's post
			ImageKey: "posts/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
		},
		{
			ID:       uuid.New().String(),
			Caption:  "Luna's New Toy",
			UserID:   users[1].ID, // Jane's post
			ImageKey: "posts/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
		},
		{
			ID:       uuid.New().String(),
			Caption:  "Buddy's Birthday Celebration",
			UserID:   users[2].ID, // Alex's post
			ImageKey: "posts/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
		},
		{
			ID:       uuid.New().String(),
			Caption:  "Coco's New Hutch",
			UserID:   users[3].ID, // Emily's post
			ImageKey: "posts/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
		},
		{
			ID:       uuid.New().String(),
			Caption:  "Rocky's First Swimming Lesson",
			UserID:   users[4].ID, // Michael's post
			ImageKey: "posts/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
		},
		{
			ID:       uuid.New().String(),
			Caption:  "Milo's Favorite Napping Spot",
			UserID:   users[0].ID, // John's second post
			ImageKey: "posts/26c4d55c-c16b-49b7-a4ef-5daa6ef2777f-BAB51C25-2C0A-4EC9-B7F5-96CAE90B0C48.jpg",
		},
	}
}

// CommentData returns a list of sample comments
func CommentData(posts []models.Post, users []models.User) []models.Comment {
	return []models.Comment{
		{
			ID:      uuid.New().String(),
			Content: "He looks so happy! What breed is he?",
			PostID:  posts[0].ID, // Comment on John's post
			UserID:  users[1].ID, // Comment by Jane
		},
		{
			ID:      uuid.New().String(),
			Content: "That's such a cute toy! Where did you get it?",
			PostID:  posts[1].ID, // Comment on Jane's post
			UserID:  users[2].ID, // Comment by Alex
		},
		{
			ID:      uuid.New().String(),
			Content: "Happy birthday, Buddy! That cake looks delicious.",
			PostID:  posts[2].ID, // Comment on Alex's post
			UserID:  users[3].ID, // Comment by Emily
		},
		{
			ID:      uuid.New().String(),
			Content: "What a nice hutch! Coco must be very happy.",
			PostID:  posts[3].ID, // Comment on Emily's post
			UserID:  users[4].ID, // Comment by Michael
		},
		{
			ID:      uuid.New().String(),
			Content: "Swimming is great exercise for dogs! Rocky looks like he's having fun.",
			PostID:  posts[4].ID, // Comment on Michael's post
			UserID:  users[0].ID, // Comment by John
		},
		{
			ID:      uuid.New().String(),
			Content: "Haha, classic cat behavior! Milo is adorable.",
			PostID:  posts[5].ID, // Comment on John's second post
			UserID:  users[1].ID, // Comment by Jane
		},
		{
			ID:      uuid.New().String(),
			Content: "I love the park too! We should arrange a playdate for our dogs.",
			PostID:  posts[0].ID, // Second comment on John's post
			UserID:  users[2].ID, // Comment by Alex
		},
	}
}

// LikeData returns a list of sample likes
func LikeData(posts []models.Post, users []models.User) []models.Like {
	return []models.Like{
		{
			PostID: posts[0].ID, // Like on John's post
			UserID: users[1].ID, // Like by Jane
		},
		{
			PostID: posts[0].ID, // Like on John's post
			UserID: users[2].ID, // Like by Alex
		},
		{
			PostID: posts[1].ID, // Like on Jane's post
			UserID: users[0].ID, // Like by John
		},
		{
			PostID: posts[1].ID, // Like on Jane's post
			UserID: users[3].ID, // Like by Emily
		},
		{
			PostID: posts[2].ID, // Like on Alex's post
			UserID: users[0].ID, // Like by John
		},
		{
			PostID: posts[2].ID, // Like on Alex's post
			UserID: users[1].ID, // Like by Jane
		},
		{
			PostID: posts[3].ID, // Like on Emily's post
			UserID: users[2].ID, // Like by Alex
		},
		{
			PostID: posts[4].ID, // Like on Michael's post
			UserID: users[3].ID, // Like by Emily
		},
		{
			PostID: posts[5].ID, // Like on John's second post
			UserID: users[1].ID, // Like by Jane
		},
		{
			PostID: posts[5].ID, // Like on John's second post
			UserID: users[4].ID, // Like by Michael
		},
	}
}

func FollowRelationData(users []models.User) []models.FollowRelation {
	return []models.FollowRelation{
		{
			FromID: users[0].ID,
			ToID:   users[1].ID,
		},
		{
			FromID: users[2].ID,
			ToID:   users[3].ID,
		},
		{
			FromID: users[4].ID,
			ToID:   users[3].ID,
		},
		{
			FromID: users[5].ID,
			ToID:   users[6].ID,
		},
		{
			FromID: users[6].ID,
			ToID:   users[5].ID,
		},
	}
}
