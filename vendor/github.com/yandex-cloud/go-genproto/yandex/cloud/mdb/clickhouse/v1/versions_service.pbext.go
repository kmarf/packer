// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

func (m *ListVersionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListVersionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListVersionsResponse) SetVersion(v []*Version) {
	m.Version = v
}

func (m *ListVersionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
