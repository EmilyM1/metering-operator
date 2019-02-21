// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReports implements ReportInterface
type FakeReports struct {
	Fake *FakeMeteringV1alpha1
	ns   string
}

var reportsResource = schema.GroupVersionResource{Group: "metering.openshift.io", Version: "v1alpha1", Resource: "reports"}

var reportsKind = schema.GroupVersionKind{Group: "metering.openshift.io", Version: "v1alpha1", Kind: "Report"}

// Get takes name of the report, and returns the corresponding report object, and an error if there is any.
func (c *FakeReports) Get(name string, options v1.GetOptions) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reportsResource, c.ns, name), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// List takes label and field selectors, and returns the list of Reports that match those selectors.
func (c *FakeReports) List(opts v1.ListOptions) (result *v1alpha1.ReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reportsResource, reportsKind, c.ns, opts), &v1alpha1.ReportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReportList{ListMeta: obj.(*v1alpha1.ReportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reports.
func (c *FakeReports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reportsResource, c.ns, opts))

}

// Create takes the representation of a report and creates it.  Returns the server's representation of the report, and an error, if there is any.
func (c *FakeReports) Create(report *v1alpha1.Report) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reportsResource, c.ns, report), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// Update takes the representation of a report and updates it. Returns the server's representation of the report, and an error, if there is any.
func (c *FakeReports) Update(report *v1alpha1.Report) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reportsResource, c.ns, report), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReports) UpdateStatus(report *v1alpha1.Report) (*v1alpha1.Report, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(reportsResource, "status", c.ns, report), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// Delete takes name of the report and deletes it. Returns an error if one occurs.
func (c *FakeReports) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reportsResource, c.ns, name), &v1alpha1.Report{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reportsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReportList{})
	return err
}

// Patch applies the patch and returns the patched report.
func (c *FakeReports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}
