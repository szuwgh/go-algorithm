package partition_labels

import (
	"fmt"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	s := "ababcbacadefegdehijhklij"
	fmt.Println("Partition lengths:", partitionLabels(s))
}
