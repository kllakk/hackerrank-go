package queue_using_two_stacks

func QueueUsingTwoStacks(queries [][]string) []string {

	var result []string

	var queue []string

	for _, query := range queries {
		switch query[0] {
			// Enqueue
			case "1":
				queue = append(queue, query[1])
				break
			// Dequeue
			case "2":
				_, queue = queue[0], queue[1:]
				break
			// Print
			case "3":
				result = append(result, queue[0])
				break
		}
	}

	return result
}