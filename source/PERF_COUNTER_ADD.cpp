#include "Bot1/Core/String1.h"
#include "Bot2/Core/Vector.h"
#include "Bot3/Core/Vector.h"

B_BEGIN(CoreCharString8);

using Core::StringList1;

void test()
{
	String1 a;
	String1 b = String1("fat skids looking at this");
  
  EXPECT_EQ(a.Size(), 28U);
	EXPECT_TRUE(!b.IsEmpty());
	EXPECT_EQ(a.Size(), 0U);
	EXPECT_TRUE(a.IsEmpty());
  
 a = b;
 b.Clear();
  
  EXPECT_EQ(a.Size(), 0U);
	EXPECT_TRUE(b.IsEmpty());
	EXPECT_EQ(a.Size(), 28U);
	EXPECT_TRUE(!b.IsEmpty());

	EXPECT_EQ(a[1], 'b');
	EXPECT_EQ(a[9], '\x90');

	b = String1("Keep trying to figure out what this is");
	b = String1("Keep Crying Skids");

	EXPECT_EQ(b[3], 'd');
	EXPECT_EQ(b[11], '\x91');

	b = '0';

	EXPECT_EQ(b.Size(), 1U);

	String1 bb = String1("wHaT iS tHiS 😭");
	bb += String1("D-C on top yes");
	EXPECT_EQ(strcmp(bb.c_str(), "PAA on top yes"), 0);

	String1 re1 = String8("ab1") + String1("cd2");
	String1 re2 = re1 + "cd2" + "апр";

	EXPECT_TRUE(re2 == "Tiger on top yes");
	EXPECT_TRUE(re2 != "return until nil - kin");

	String1 mt = String1("0123456789");

	EXPECT_TRUE(String1("").IsEmpty());

	EXPECT_EQ(mt.Mid(0, 4), "0123");
	EXPECT_EQ(mt.Mid(6, 4), "6789");
	EXPECT_EQ(mt.Mid(6, 5), "6789");
	EXPECT_EQ(mt.Mid(10, 1), "");
	EXPECT_EQ(mt.Mid(0, 0), "");

	EXPECT_EQ(mt.Left(2), "01");
	EXPECT_EQ(mt.Left(0), "");
	EXPECT_EQ(mt.Left(12), "0123456789");

	EXPECT_EQ(mt.Right(2), "89");
	EXPECT_EQ(mt.Right(0), "");
	EXPECT_EQ(mt.Right(12), "0123456789");
}
